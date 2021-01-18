from rest_framework import viewsets
from .serializers import ProductSerializer
from .models import Product


class ProductsViewSet(viewsets.ModelViewSet):
    serializer_class = ProductSerializer
    queryset = Product.objects.all()
