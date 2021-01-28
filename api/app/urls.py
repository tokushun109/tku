from django.urls import path
from .views import ProductListAPIView, AccessoryListAPIView, ProductRetrieveAPIView

urlpatterns = [
    path('product/', ProductListAPIView.as_view()),
    path('product/<pk>/', ProductRetrieveAPIView.as_view()),
    path('accessory_category/', AccessoryListAPIView.as_view()),
    ]
