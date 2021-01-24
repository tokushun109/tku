from django.contrib import admin
from .models import Product, AccessoryCategory

admin_model_list = [Product, AccessoryCategory]
for admin_model in admin_model_list:
    admin.site.register(admin_model)
