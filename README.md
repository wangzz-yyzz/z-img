## Z-IMG

### introduction
Z-IMG is a simple image server, it can be used to store and display the image.


### api(default_port:8086)
#### get image
>url: /img/get  
method: get  
param: id  
return: image

#### post image
>url: /img/post  
method: post  
param: form.file  
return: json{id: "image id"}

#### delete image
>url: /img/delete  
method: get  
param: id  
return: json{msg: "result"}



### deploy
```
git clone
cd z-img
docker build -t z-img .
docker run -v /Users/wzz/docker_file/z-img/img:/go/src/wzz/z-img/img --name z-img -p 8086:8086 -d z-img
```

### test
open the browser and input the url:
http://127.0.0.1:8086/ui/index  
if you can see the image, it's ok.
