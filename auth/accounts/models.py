from django.contrib.auth.models import AbstractUser
from django.db import models

from shared.models import BaseModel

# Create your models here.


class User(AbstractUser, BaseModel):
    full_name = models.CharField(max_length=255, blank=True, null=True)
    test = models.IntegerField(default=0)
