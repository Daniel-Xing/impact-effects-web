FROM python:3

WORKDIR /usr/src/app
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple/
RUN python -m pip install --upgrade pip
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt 
RUN pip config set global.index-url https://pypi.python.org/simple/
# RUN pip install ImpactEffect==0.0.7
COPY ./dist/ImpactEffect-0.0.7-py3-none-any.whl .
# RUN cd ImpactEffect && pip install -e . && cd ..
RUN pip install ImpactEffect-0.0.7-py3-none-any.whl
COPY . .
EXPOSE 50051
CMD [ "python", "./service.py" ]