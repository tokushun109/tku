from rest_framework import serializers
from app.models.producer_profile import ProducerProfile


class ProducerProfileSerializer(serializers.ModelSerializer):
    class Meta:
        model = ProducerProfile
        fields = ('name', 'introduction', 'logo')
