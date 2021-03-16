from django.urls import path
from app.views import ProductListAPIView, ProductRetrieveAPIView, AccessoryCategoryListAPIView, MaterialCategoryListAPIView, \
    SalesSiteListAPIView, ProducerProfileAPIView, SkillMarketListAPIView, SnsListAPIView

urlpatterns = [
    path('product/', ProductListAPIView.as_view()),
    path('product/<pk>/', ProductRetrieveAPIView.as_view()),
    path('accessory_category/', AccessoryCategoryListAPIView.as_view()),
    path('material_category/', MaterialCategoryListAPIView.as_view()),
    path('sales_site/', SalesSiteListAPIView.as_view()),
    path('producer_profile/', ProducerProfileAPIView.as_view()),
    path('skill_market/', SkillMarketListAPIView.as_view()),
    path('sns/', SnsListAPIView.as_view()),
]
