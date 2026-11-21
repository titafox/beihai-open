

function request (options) {
	//const baseURL  = 'http://<api-host>:8000/v1/';
	 const baseURL  = '/v1/';
	//
	let language = getCookie("userLanguage");

	if(language){
		language = language.slice(0,1).toUpperCase()+language.slice(1);
	} else {
		language = 'En'
	}

	let token = getCookie('token');
	let chainId = getCookie('chainId')
	console.log(token)
	let headers = {
		"language":language,

	}
	console.log(chainId)
	if(chainId) {
		headers['Chainid'] = chainId;
	}
	if(token){
		token = `Bearer ${token}`;
		headers['Authorization'] = token;
	}

	let method  = options.method;
	if(!method) {
		method = 'get';
	} else {
		headers['Content-Type'] = 'application/json'
	}
	// let async = options.async;
	// if(!async) method = false;

	// console.log('method',method)
	let url = options.url;
	let data = options.data;

	if(method == 'post')
		data = JSON.stringify(data)

	return new Promise(function ( resolve,reject ) {



		$.ajax({
			method: method,
			url: baseURL + options.url,
			data: data,
			async: true,
			dataType: 'json',
			headers: headers,
			success:(res)=>{
				if(res.code == 0){
					resolve( res.data )
				} else {
					console.log(err)
				}
			},
			error:(XmlHttpRequest,textStatus, errorThrown)=>{
				console.log(XmlHttpRequest)
				if(XmlHttpRequest.responseJSON?.code == '403'){
					//获取token
					if (metamask && typeof metamask.isInstall == 'function' && metamask.isInstall()){
						metamask.getTokenInfo()
					}
				} else {

				}
				reject( textStatus,errorThrown )
			},
		})

		// axios({
		//     method: method,
		//     url: url,
		//     data: data
		// }).then( res => resolve( res )).catch( err => reject( err ))
		// $.ajax({
		// 	method: method,
		// 	url: url,
		// 	async: false
		// })


	})
}

