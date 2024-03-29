const HOST="http://0.0.0.0:50052";
//upload image files(*.png,*.jpe?g)
const AppendFileAPI = HOST + "/api/appendFile";
//upload compressed file(*.zip)
const CreateDatasetAPI = HOST + "/api/createDataset";
const NAME = "Imapct Earth !!!";
function ajaxPost(path, body, noCode) {
  // let token = localStorage.getItem("token");
  let headers = {
    "Content-Type": "application/json"
  };
  console.log('body' + JSON.stringify(body))
  console.log('path' + path)
  return new Promise((resolve, reject) => {
    fetch(HOST + path, {
      method: "POST",
      body:JSON.stringify(body)
    })
      .then(response => {
        console.log('response' + response)
        return response.json();
      })
      .then(data => {
        console.log('data' + data)
        if (data.code != 200 && !noCode) {
          // reject({code:data.code,message:data.message});
          reject(data)
        } else {
          resolve(data);
        }
      })
      .catch((response, error) => {
        console.log('error' + error);
        reject({ code: -1, message: "登陆失败" });
      });
  });
}
function ajaxFile(path, body) {
  let token = localStorage.getItem("token");
  let headers = {
    "Content-Type": "application/json"
  };
  headers["Authorization"] = "Bearer " + (token ? token : "");
  return new Promise((resolve, reject) => {
    fetch(HOST + path, {
      method: "POST",
      credentials: "include",
      changeOrigin:true,
      headers,
      body: JSON.stringify(body)
    })
      .then(response => {
        if (response.headers.Authorization && response.headers.Authorization.replace("Bearer ", "") != token) {
          localStorage.setItem("token", response.headers.Authorization.replace("Bearer ", ""));
        }
        if (response.headers.get("content-type") == "application/octet-stream") {
          let f = "";
          let j = response.json();
          j.then(data => {
            let d = { f: f, d: data };
            resolve(d);
          });
        } else {
          let d = response.json();
          d.then(data => {
            reject({ code: data.code, message: data.message });
          });
        }
      })
      .catch((response, error) => {
        console.log(error);
        reject({ code: -1, message: "未知错误" });
      });
  });
}
function ajax(path) {
  let token = localStorage.getItem("token");
  return new Promise((resolve, reject) => {
    fetch(HOST + path, {
      credentials: "include",
      headers: {
        Authorization: "Bearer " + (token ? token : "")
      }
    })
      .then(response => {
        if (response.headers.Authorization && response.headers.Authorization.replace("Bearer ", "") != token) {
          localStorage.setItem("token", response.headers.Authorization.replace("Bearer ", ""));
        }
        return response.json();
      })
      .then(data => {
        if (data.code == 200) {
          resolve(data);
        } else {
          reject(data);
        }
      })
      .catch(() => {
        reject({ code: -1, message: "未知错误" });
      });
  });
}
function ajaxImg(path) {
  return new Promise((resolve, reject) => {
    // fetch( path)
    //   .then(response => {
    //     return response.blob();
    //   })
    //   .then(data => {
    //     var reader = new FileReader();
    //     reader.onload = function() {
    //       resolve(this.result);
    //     };
    //     reader.readAsDataURL(data);
    //   })
    //   .catch(() => {
    //     reject({ code: -1, message: "未知错误" });
    //   });

    var e = new Image();
    e.addEventListener('load', function() {
      console.log(e);
    });
    e.src = path;
    // var reader = new FileReader();
    // reader.onload = function() {
    //   resolve(this.result);
    // };
    // reader.readAsDataURL(path);
  });
}
function checkIsLogin(path) {
  let token = localStorage.getItem("token");
  console.log(token);
  return new Promise((resolve, reject) => {
    fetch(HOST + path, {
      credentials: "include",
      method: 'POST',
      headers: {
        Authorization: "Bearer " + (token ? token : "")
      }
    })
      .then(response => {
        if (response.headers.get('Authorization') &&response.headers.get('Authorization').replace("Bearer ", "") != token) {
          localStorage.setItem("token", response.headers.get('Authorization').replace("Bearer ", ""));
        }
        if (response.status == 302 || response.status == 500) {
          reject();
          return;
        }
        return response.json();
      })
      .then(data => {
        if (data.code == 500) {
          reject();
          return;
        }
        resolve();
      })
      .catch(() => {
        reject();
      });
  });
}
function getUserGroupStr(g) {
  let group = parseInt(g);
  return group == 1 ? "user" : group == 3 ? "admin" : "reviewer";
}
export { HOST, AppendFileAPI, ajax, ajaxPost, ajaxFile, CreateDatasetAPI, getUserGroupStr, NAME, checkIsLogin, ajaxImg };
