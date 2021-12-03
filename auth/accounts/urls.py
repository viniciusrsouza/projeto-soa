from django.conf.urls import url
from django.urls.conf import include, re_path
from rest_framework_jwt.views import obtain_jwt_token, refresh_jwt_token, verify_jwt_token
from rest_framework_nested import routers
from accounts import views

user_router = routers.SimpleRouter()
user_router.register(r'users', views.UserViewSet, basename='user')

urlpatterns = [
    url(r'^token/', obtain_jwt_token),
    url(r'^refresh-token/', refresh_jwt_token),
    url(r'^introspection/', verify_jwt_token),
    re_path('^', include(user_router.urls)),
]
