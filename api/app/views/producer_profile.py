from django.shortcuts import get_object_or_404
from rest_framework import status, views
from rest_framework.response import Response

from app.models.producer_profile import ProducerProfile
from app.serializers.producer_profile import ProducerProfileSerializer


class ProducerProfileAPIView(views.APIView):
    '''
    製作者モデルの取得APIクラス
    '''

    def get(self, request, *args, **kwargs):
        # pk=1の製作者のプロフィールを取得する
        producer_profile = get_object_or_404(ProducerProfile, pk=1)
        serializer = ProducerProfileSerializer(
            producer_profile, context={"request": request})
        return Response(serializer.data, status.HTTP_200_OK)
