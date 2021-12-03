from django.contrib import admin
from django.contrib.auth.admin import UserAdmin as BaseUserAdmin
from . import models

# Register your models here.


@admin.register(models.User)
class UserAdmin(BaseUserAdmin):
    list_display = ('id', 'email', 'username', 'full_name',
                    'last_login',  'date_joined')
    search_fields = ('email', 'username', 'full_name', 'id')

    add_fieldsets = (
        (
            None,
            {
                "classes": ("wide",),
                "fields": ("username", "full_name", ("password1", "password2"), "email"),
            },
        ),
        (
            "Permissions",
            {
                "fields": ("is_active", "is_staff", "is_superuser", "groups", "user_permissions"),
            },
        ),
    )
