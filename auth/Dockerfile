FROM python:3.8.12

ENV PYTHONUNBUFFERED 1

RUN mkdir /app
WORKDIR /app

COPY ./requirements/ /app/requirements/
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r /app/requirements/base.txt

COPY . /app/