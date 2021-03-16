from rest_framework import generics

from app.models.site import SalesSite, SkillMarket, Sns
from app.serializers.site import SalesSiteSerializer, SkillMarketSerializer, SnsSerializer


class SalesSiteListAPIView(generics.ListAPIView):
    '''
    販売サイトモデルの取得(一覧)APIクラス
    '''
    queryset = SalesSite.objects.all()
    serializer_class = SalesSiteSerializer


class SkillMarketListAPIView(generics.ListAPIView):
    '''
    販売サイトモデルの取得(一覧)APIクラス
    '''
    queryset = SkillMarket.objects.all()
    serializer_class = SkillMarketSerializer


class SnsListAPIView(generics.ListAPIView):
    '''
    販売サイトモデルの取得(一覧)APIクラス
    '''
    queryset = Sns.objects.all()
    serializer_class = SnsSerializer
