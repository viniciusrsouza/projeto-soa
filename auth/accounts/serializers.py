from rest_framework import serializers
from rest_framework.generics import get_object_or_404
from rest_framework_simplejwt.tokens import UntypedToken
from accounts.models import User


class RetrieveUserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ('id', 'username', 'email')


class CreateUserSerializer(serializers.ModelSerializer):
    confirm_password = serializers.CharField(write_only=True)

    def validate(self, attrs):
        print(attrs)
        if attrs.get('password') != attrs.pop('confirm_password'):
            raise serializers.ValidationError('Passwords do not match')
        return super().validate(attrs)

    class Meta:
        model = User
        fields = ('id', 'full_name', 'username', 'email',
                  'password', 'confirm_password')
        extra_kwargs = {"full_name": {"required": True},
                        "email": {"required": True}}


class UpdateUserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ('id', 'username', 'email', 'full_name')
        read_only_fields = ('id', 'username', 'email')


class TokenIntrospectionSerializer(serializers.Serializer):
    token = serializers.CharField()

    def validate(self, attrs):
        token = UntypedToken(attrs['token'])
        user_id = token.get('user_id')
        if not user_id:
            raise serializers.ValidationError('Invalid token')

        user = get_object_or_404(User, id=user_id)
        serializer = RetrieveUserSerializer(user)
        return serializer.data
