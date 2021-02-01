from django.contrib import admin
from .models import Product, AccessoryCategory, MaterialCategory, SalesSite, ProductImage, ProducerProfile, SkillMarket, Sns

admin_model_list = [
    Product,
    AccessoryCategory,
    MaterialCategory,
    SalesSite,
    ProductImage,
    ProducerProfile,
    SkillMarket,
    Sns,
    ]

for admin_model in admin_model_list:
    admin.site.register(admin_model)
