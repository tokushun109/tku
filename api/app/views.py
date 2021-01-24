from rest_framework import viewsets
from .serializers import ProductSerializer, AccessoryCategorySerializer
from .models import Product, AccessoryCategory


class ProductsViewSet(viewsets.ModelViewSet):
    serializer_class = ProductSerializer
    queryset = Product.objects.all()


class AccessoryCategoryViewSet(viewsets.ModelViewSet):
    serializer_class = AccessoryCategorySerializer
    queryset = AccessoryCategory.objects.all()
