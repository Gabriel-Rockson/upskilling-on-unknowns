FROM python:3.12.2-bullseye

WORKDIR /usr/src/app

COPY Pipfile ./
COPY Pipfile.lock ./

RUN pip install pipenv

RUN pipenv install --system

COPY . .

EXPOSE 2222

CMD [ "python", "main.py" ]
