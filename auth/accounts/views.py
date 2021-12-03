from calendar import c
from rest_framework import (viewsets, generics, permissions)
from rest_framework_simplejwt.views import TokenViewBase
from accounts.models import User
from accounts.serializers import RetrieveUserSerializer, TokenIntrospectionSerializer, CreateUserSerializer, UpdateUserSerializer

# Create your views here.


class UserViewSet(viewsets.ModelViewSet):
    serializers = {
        'default': RetrieveUserSerializer,
        'create': CreateUserSerializer,
        'update': UpdateUserSerializer,
        'partial_update': UpdateUserSerializer,
        'list': RetrieveUserSerializer,
    }

    def create(self, request, *args, **kwargs):
        if "username" in self.request.data:
            self.request.data["email"] = self.request.data["username"]
        return super().create(request, *args, **kwargs)

    def get_serializer_context(self):
        return {"request": self.request}

    def get_serializer_class(self):
        return self.serializers.get(self.action, self.serializers.get("default"))

    def get_queryset(self):
        return User.objects.all()


class RetrieveUserApiView(generics.RetrieveAPIView):
    queryset = User.objects.all()
    serializer_class = RetrieveUserSerializer
    permission_classes = [permissions.IsAuthenticated]

    def get_object(self):
        return self.request.user


class TokenIntrospectionView(TokenViewBase):
    serializer_class = TokenIntrospectionSerializer
