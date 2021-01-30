from django.urls import path
from .views import ProductListAPIView, ProductRetrieveAPIView, AccessoryCategoryListAPIView, MaterialCategoryListAPIView

urlpatterns = [
    path('product/', ProductListAPIView.as_view()),
    path('product/<pk>/', ProductRetrieveAPIView.as_view()),
    path('accessory_category/', AccessoryCategoryListAPIView.as_view()),
    path('material_category/', MaterialCategoryListAPIView.as_view()),
    ]
