
let enter = () => {
    let pwd = $('#password').val();
    $.ajax({
        url: '/issues/secret',
        method:'POST',
        dataType: 'json',
        contentType: 'application/json; charset=utf-8',
        data:JSON.stringify({secret:pwd}),
        success:(xhr,status,msg)=>{
            console.log(msg);
        }
    });
};

$('#login').click(enter);
