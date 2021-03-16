from rest_framework import generics

from app.models.category import AccessoryCategory, MaterialCategory
from app.serializers.category import AccessoryCategorySerializer, MaterialCategorySerializer


class AccessoryCategoryListAPIView(generics.ListAPIView):
    '''
    アクセサリーカテゴリーモデルの取得(一覧)APIクラス
    '''
    queryset = AccessoryCategory.objects.all()
    serializer_class = AccessoryCategorySerializer


class MaterialCategoryListAPIView(generics.ListAPIView):
    '''
    材料カテゴリーモデルの取得(一覧)APIクラス
    '''
    queryset = MaterialCategory.objects.all()
    serializer_class = MaterialCategorySerializer
