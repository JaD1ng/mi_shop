(function ($) {
    var app = {
        init: function () {
            this.addAddress();
            this.changeDefaultAddress();
            this.editAddress();

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

                $.post('/address/addAddress', {name: name, phone: phone, address: address}, function (response) {
                    console.log(response)

                    if (response.success) {
                        var addressList = response.result;
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
            $(".addressModify").click(function () {

                //请求接口获取当前收货地址id对应的数据
                var addressId = $(this).parent().attr("data-id")
                $.get("/address/getOneAddressList", {"addressId": addressId}, function (response) {
                    console.log(response)
                    if (response.success) {
                        var addressInfo = response.result;
                        $("#edit_id").val(addressInfo.id);
                        $('#edit_name').val(addressInfo.name);
                        $('#edit_phone').val(addressInfo.phone);
                        $('#edit_address').val(addressInfo.address);
                    } else {
                        alert(response.message)
                    }
                    $('#editAddressModal').modal('show')
                })

            })

            $("#editAddress").click(function () {
                var id = $('#edit_id').val();
                var name = $('#edit_name').val();
                var phone = $('#edit_phone').val();
                var address = $('#edit_address').val();
                if (name == '' || phone == "" || address == "") {
                    alert('姓名、电话、地址不能为空')
                    return false;
                }
                var reg = /^[\d]{11}$/;
                if (!reg.test(phone)) {
                    alert('手机号格式不正确');
                    return false;
                }
                $.post('/address/editAddress', {
                    id: id, name: name, phone: phone, address: address
                }, function (response) {
                    if (response.success) {
                        var addressList = response.result;
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

                    $('#editAddressModal').modal('hide')
                });
            })
        },
    }
    $(function () {
        app.init();
    })
})($)
