$(function(){
	const language = getCookie('userLanguage');
	const languageArr = ['en','vi','cn'];
	const langNameArr = ['English','Tiếng việt','中文'];
	const langName = langNameArr[languageArr.indexOf(language)]
	$('.footer__lang-btn').find('span').text(langName)
	$('.footer__lang-dropdown').find('li').on('click',function(){
		var language =$(this).attr('name');
		console.log(language);
		getCookie("userLanguage", language, {
			expires: 30,
			path: '/'
		});
		location.reload();
	})
})