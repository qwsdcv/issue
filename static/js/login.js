
let enter = () => {
    let pwd = $('#password').val();
    $.ajax({
        url: '/issues/secret',
        method:'POST',
        contentType: 'application/json; charset=utf-8',
        data:JSON.stringify({secret:pwd}),
        success:(msg,status,xhr)=>{
            console.log(msg);
            let token = xhr.getResponseHeader('Authorization');
            $.ajaxSetup({
                headers: {
                    "Authorization": token
                }
            });
            window.location = msg.target+"?token="+token;
        }
    });
};

$('#login').click(enter);
