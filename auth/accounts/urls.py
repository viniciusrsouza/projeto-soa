from django.conf.urls import url
from django.urls.conf import include, re_path
from rest_framework_simplejwt import views as jwt_views
from rest_framework_nested import routers
from accounts import views


user_router = routers.SimpleRouter()
user_router.register('users', views.UserViewSet, basename='user')

urlpatterns = [
    url(r'^token/$', jwt_views.TokenObtainPairView.as_view(),
        name='token_obtain_pair'),
    url(r'^refresh-token/$', jwt_views.TokenRefreshView.as_view(), name='token_refresh'),
    url(r'^introspection/$', views.TokenIntrospectionView.as_view(),
        name='token_verify'),
    url(r'^user/$', views.RetrieveUserApiView.as_view()),
    re_path('^', include(user_router.urls)),
]
