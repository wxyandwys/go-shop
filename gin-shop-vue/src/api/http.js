import axios from "axios";
import qs from "qs";
import store from './../store/index'
import {
    Dialog
} from "vant";

// axios.defaults.baseURL = 'http://hp.suoluomei.cn/index.php'  //正式
// axios.defaults.baseURL = 'http://hp.suoluomei.cn/index.php' //测试
axios.defaults.baseURL = 'http://localhost:8901/v1'

//post请求头
axios.defaults.headers.post["Content-Type"] =
    "application/x-www-form-urlencoded;charset=UTF-8";

axios.defaults.transformRequest = [obj => qs.stringify(obj)] 
//设置超时
axios.defaults.timeout = 20000;

axios.interceptors.request.use(
    config => {
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

axios.interceptors.response.use(
    response => {
        if (response.status == 200) {
            return Promise.resolve(response);
        } else {
            return Promise.reject(response);
        }
    },
    error => {
        if (error.response.status == 402) {
            Dialog.alert({
                title: "提示",
                message: error.response.data.text//"用户超时，请重新登录"
            });
            store.state.user = null
        } else {
            Dialog.alert({
                title: "提示",
                message: "网络请求失败，请刷新重试"
            });
        }
        
    }
);
export default {
    post(url, data) {
        return new Promise((resolve, reject) => {
            axios({
                method: 'post',
                url,
                data: data,
                headers: {
                    'Token': 11111,
                    'AccessToken': store.state.user != null ? store.state.user : ''
                }
            })
                .then(res => {
                    resolve(res.data)
                })
                .catch(err => {
                    reject(err)
                });
        })
    },

    get(url, data) {
        return new Promise((resolve, reject) => {
            axios({
                method: 'get',
                url,
                params: data,
                headers: {
                    'Token': 11111
                }
            })
                .then(res => {
                    resolve(res.data)
                })
                .catch(err => {
                    reject(err)
                })
        })
    }
};
