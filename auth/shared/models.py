from django.db import models

# Create your models here.

class BaseModel(models.Model):
  class Meta:
    abstract = True
  
  created_at = models.DateTimeField(auto_now_add=True)
  updated_at = models.DateTimeField(auto_now=True)

  def save(self, *args, **kwargs):
    self.updated_at = models.DateTimeField(auto_now=True)
    super().save(*args, **kwargs)