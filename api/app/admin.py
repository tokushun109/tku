from django.contrib import admin
from .models import Product, AccessoryCategory, MaterialCategory, SalesSite, ProductImage

admin_model_list = [Product, AccessoryCategory, MaterialCategory, SalesSite, ProductImage]
for admin_model in admin_model_list:
    admin.site.register(admin_model)
