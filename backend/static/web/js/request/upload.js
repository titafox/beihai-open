
const API = {
    getUploadParams: 'api/credential', // 获取上传的云厂商参数
    reportUploadResult: 'api/v1/upload-result', // 上传完成后，将上传结果返回给后端
};

(function (global, undefined) {


    //定义了一个名字为Cal的类
    class Upload {
        //constructor是一个构造方法，用来接收参数
        constructor(options) {
			console.log(options)
            //检测传入的参数类型是否为对象，若不是则抛出错误
            if (arguments.length > 0 && (typeof options).toLowerCase() !== 'object') {
                throw new TypeError(options + 'is not a Object');
            }
            //传入参数覆盖默认值
            //this代表的是实例对象
            this.options = Object.assign({
                mediaType: 'image',
                sceneName: '',
                businessName: 'test',
                mode: 'dev',
                limit: 1,
                token: '',
                getQr: false,
                maxSize: 50 * 1024 * 1024,
                startUpload: () => { },
                imageKey: '',
                cos: null

            }, options);
            this.init();
        }
		
		
		

       
        //这是一个类的方法，注意千万不要加上function
        init() {
            console.log(`你输入的参数为${JSON.stringify(this.options)}`);
        }
        //方法之间不要用逗号分隔，否则会报错

        //重置数据
        resetData() {

        }
        //获取上传的云厂商参数
        setCos(){
            return new Promise((resolve, reject) => {
                const { mediaType, sceneName } = this.options;
                request({
                    url: API.getUploadParams,
                    data: {
                        mediaType,
                        sceneName,
                        businessName: 'test',
                        source: 'web',
                    },
                }).then(res=>{
                    this.cos = new COS({
                        getAuthorization: (options, callback) => {
                            callback({
                                TmpSecretId: res.tmpSecretID,
                                TmpSecretKey: res.tmpSecretKey,
                                XCosSecurityToken: res.sessionToken,
                                ExpiredTime: ((new Date).valueOf() / 1000) | 0 + 30 * 24 * 60 * 60,
                            });
                        },
                    });
                    resolve(res);
                }).catch(err=>{
					showToast({
						message: err,
						type:'failed'
					})
				})
            });
        }

        upload = (files, callback) => {
            this.selectFileCallback = callback
            return new Promise((resolve, reject) => {
                this.resolve = (res) => {
                    this.resetData();
                    resolve(res);
                };
                this.reject = (err) => {
                    this.resetData();
                    reject(err);
                };

                if (files && files.length > 0) {
                    this.startUpload(files);
                } else {
                    this.createFrom();
                }
            })
        }

        createFrom () {
            const PC = !!navigator.platform.match(/mac|win/i);
            this.input = document.createElement('input');
            this.input.type = 'file';
            this.input.accept = PC ? 'image/jpeg, image/png, image/jpg, image/gif' : 'image/*';
            this.input.name = 'FileContent';
            // 支持多张上传
            this.input.multiple = this.options.multiple;//false;
            this.input.style.cssText = 'position: fixed; height:1em; width:1em; opacity:1 ; top:-100em; left:0;';
            this.input.id = 'wpt_uploadimg_input_id';
            document.body.appendChild(this.input);
            this.input.onchange = () => {
                this.startUpload();
            };
            this.input.click();
        }

        startUpload (files)  {
			return new Promise((resolve, reject) => {
				if (!files) files = this.input.files || [];
				if (files.length === 0) {
					showToast({
						message:$.i18n.prop('create-chooseImg'),
						type:'failed'
					})
				    return;
				}
				if(files.length > this.options.limit){
				  
					showToast({
						message: $.i18n.prop('create-imgLimitLength'),
						type:'failed'
					})
				    return;
				}
				const limitFiles = [...files].slice(0, Math.min(this.options.limit || 1, files?.length || 1));
				const execUpload = [...limitFiles].some((file) => {
					console.log('size')
				    if (file.size > this.options.maxSize) {
						showToast({
							message: $.i18n.prop('create-imgLimitSize'),
							type:'failed'
						})
				        return true;
				    }
				    return false;
				});
				if (execUpload) return;
				this.options.startUpload?.(files.length);
				this.selectFileCallback && this.selectFileCallback();
				
				this.afterStartUpload(limitFiles)
			   
			})
			
			
        }
		afterStartUpload = async(limitFiles) =>{
			try {
			    await this.setCos();
			    const res = await Promise.all(limitFiles.map((file, index) => this.uploadImgObject(file, index)));
			    this.uploadSuccess(res, limitFiles);
			} catch (error) {
			    this.reject(error);
			}
		}


        uploadSuccess = (imgsInfo) => {
            const serverIds = [];
            imgsInfo.forEach((item) => {
                const url = item.Location.replace('coltstail-1251022884.cos.ap-singapore.myqcloud.com', 'coltstail-1251022884.file.myqcloud.com');
                serverIds.push(url.startsWith('http') ? url : `https://${url}`);
            });

            let result = {
                serverIds,
            };
            this.resolve && this.resolve(result);
        };

        // 调用腾讯云api上传
        uploadImgObject (file, index) {
            return new Promise((resolve, reject) => {
                // const { region, bucket } = state;
				const region = 'ap-singapore';
				const bucket = 'coltstail-1251022884';
                const { imageKey } = this.options;
                const uuid = this.uuid;
                const key = [8, 4, 4, 4, 12].map(uuid).join('-').toLowerCase();
                this.cos.putObject(
                    {
                        Bucket: bucket /* 必须 */,
                        Region: region /* 必须 */,
                        Key: imageKey || key /* 必须 */,
                        Body: file, // 上传文件对象
                        onProgress: (info) => {
                            this.options.onProgress && this.options.onProgress(info, index);
                        },
                    },
                    (err, data) => {
                        if (err) {
                            // WPT.Modal.loading('error: 腾讯云上传失败', { time: 2000 });
                            // 上传失败的原因一般是签名过期了
                            return reject(err);
                        }
                        resolve({
                            ...data,
                        });
                    }
                );
            })
        }

        uuid (len = 8) {
            const S = 'qwertyuioopasdfghjklzxcvbnmQWERTYUIOOPASDFGHJKLZXCVBNM0123456789';
            const LEN = S.length - 1;
            return ' '.repeat(len).split('').map(() => S[Math.round(Math.random() * LEN)]).join('');
        }
    }

    
    global.Upload = Upload;
})(this);

// var t = new Upload({
//     x: 5,
//     y: 6
// });
// console.log(t.add());  //计算结果为11
// console.log(t.sub());  //计算结果为-1
// console.log(t.rid());  //计算结果为30
// console.log(t.div());  //计算结果为0.8333333333333334
