from django.contrib import admin
from .models import Product, AccessoryCategory, MaterialCategory

admin_model_list = [Product, AccessoryCategory, MaterialCategory]
for admin_model in admin_model_list:
    admin.site.register(admin_model)
