from rest_framework import (viewsets)
from accounts.models import User

# Create your views here.


class UserViewSet(viewsets.ModelViewSet):
    class Meta:
        model = User
