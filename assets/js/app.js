$(document).ready(function () {
    // serialize form data into json
    $.fn.serializeObject = function () {
        var o = {};
        var a = this.serializeArray();
        $.each(a, function () {
            if (o[this.name]) {
                if (!o[this.name].push) {
                    o[this.name] = [o[this.name]];
                }
                o[this.name].push(this.value || '');
            } else {
                o[this.name] = this.value || '';
            }
        });
        return o;
    };
    $(function () {
        $('#book').on('submit', function (e) {
            var url = '/api/book';
            e.preventDefault();
            var formData = $(this).serializeObject();
            fetch(url, {
                method: 'POST',
                body: JSON.stringify(formData),
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4iLCJyb2xlIjoiYWRtaW4ifQ.0U-kwiJkTvcwr56zE3_yyy1oSJx7a7G8410ECBlOApw'
                }
            }).then(res => res.json())
                .catch(error => console.error('Error:', error))
                .then(response => console.log('Success:', response));
        });
    });
});