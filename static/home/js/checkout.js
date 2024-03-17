(function ($) {
    var app = {
        init: function () {
            this.addAddress();
            this.changeDefaultAddress();
            this.editAddress();
            this.doEditAddress();

        }, addAddress: function () {
            $("#addAddress").click(function () {
                var name = $('#add_name').val();
                var phone = $('#add_phone').val();
                var address = $('#add_address').val();
                if (name == '' || phone == "" || address == "") {
                    alert('姓名、电话、地址不能为空')
                    return false;
                }
                var reg = /^[\d]{11}$/;
                if (!reg.test(phone)) {
                    alert('手机号格式不正确');
                    return false;
                }
                console.log({name: name, phone: phone, address: address})
                $.post('/address/addAddress', {name: name, phone: phone, address: address}, function (response) {
                    console.log(response)

                    if (response.success) {
                        var addressList = response.addressList;
                        var str = ""
                        for (var i = 0; i < addressList.length; i++) {
                            if (addressList[i].default_address) {
                                str += '<div class="address-item J_addressItem selected" data-id="' + addressList[i].id + '" data-name="' + addressList[i].name + '" data-phone="' + addressList[i].phone + '" data-address="' + addressList[i].address + '" > ';
                                str += '<dl>';
                                str += '<dt> <em class="uname">' + addressList[i].name + '</em> </dt>';
                                str += '<dd class="utel">' + addressList[i].phone + '</dd>';
                                str += '<dd class="uaddress">' + addressList[i].address + '</dd>';
                                str += '</dl>';
                                str += '<div class="actions">';
                                str += '<a href="javascript:void(0);" data-id="' + addressList[i].id + '" class="modify addressModify">修改</a>';
                                str += '</div>';
                                str += '</div>';

                            } else {
                                str += '<div class="address-item J_addressItem" data-id="' + addressList[i].id + '" data-name="' + addressList[i].name + '" data-phone="' + addressList[i].phone + '" data-address="' + addressList[i].address + '" > ';
                                str += '<dl>';
                                str += '<dt> <em class="uname">' + addressList[i].name + '</em> </dt>';
                                str += '<dd class="utel">' + addressList[i].phone + '</dd>';
                                str += '<dd class="uaddress">' + addressList[i].address + '</dd>';
                                str += '</dl>';
                                str += '<div class="actions">';
                                str += '<a href="javascript:void(0);" data-id="' + addressList[i].id + '" class="modify addressModify">修改</a>';
                                str += '</div>';
                                str += '</div>';
                            }
                        }

                        $("#addressList").html(str)

                    } else {
                        alert(response.message)
                    }

                    $('#addAddressModal').modal('hide')
                });

            })
        }, changeDefaultAddress: function () {


        }, editAddress: function () {


        },

        doEditAddress: function () {


        }
    }
    $(function () {
        app.init();
    })
})($)
