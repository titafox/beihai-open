define(['jquery', 'bootstrap', 'backend', 'table', 'form'], function ($, undefined, Backend, Table, Form) {

    var Controller = {
        index: function () {
            // 初始化表格参数配置
            Table.api.init({
                extend: {
                    index_url: 'tbl/account/index' + location.search,
                    add_url: 'tbl/account/add',
                    edit_url: 'tbl/account/edit',
                    del_url: 'tbl/account/del',
                    multi_url: 'tbl/account/multi',
                    import_url: 'tbl/account/import',
                    table: 'tbl_account',
                }
            });

            var table = $("#table");

            // 初始化表格
            table.bootstrapTable({
                url: $.fn.bootstrapTable.defaults.extend.index_url,
                pk: 'id',
                sortName: 'id',
                fixedColumns: true,
                fixedRightNumber: 1,
                columns: [
                    [
                        {checkbox: true},
                        {field: 'id', title: __('Id')},
                      //  {field: 'status', title: __('Status'), searchList: {"1":__('Status 1'),"0":__('Status 0')}, formatter: Table.api.formatter.status},
                      {field: 'nickname', title: __('Nickname'), operate: 'LIKE'},
                      {field: 'avatar', title: __('Avatar'), operate: 'LIKE', events: Table.api.events.image, formatter: Table.api.formatter.image},
                      {field: 'phone', title: __('Phone'), operate: 'LIKE'},
                        {field: 'bsn_address', title: __('Bsn_address'), operate: 'LIKE'},
                        {field: 'name', title: __('Name'), operate: 'LIKE'},
                        {field: 'id_card_num', title: __('Id_card_num'), operate: 'LIKE'},
                      {field: 'is_handsel', title: __('Is_handsel'), searchList: {"0":__('Is_handsel 0'),"1":__('Is_handsel 1')}, formatter: Table.api.formatter.normal},
                      //{field: 'pid', title: __('Pid')},
                      {field: 'nums', title: __('真实邀请人数'), operate: '>='},
                      {field: 'create_time', title: __('Create_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                        {field: 'update_time', title: __('Update_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                       // {field: 'authentication', title: __('Authentication')},
                        {field: 'operate', title: __('Operate'), table: table, events: Table.api.events.operate, formatter: Table.api.formatter.operate}
                    ]
                ]
            });

            // 为表格绑定事件
            Table.api.bindevent(table);
        },
        add: function () {
            Controller.api.bindevent();
        },
        edit: function () {
            Controller.api.bindevent();
        },
        api: {
            bindevent: function () {
                Form.api.bindevent($("form[role=form]"));
            }
        }
    };
    return Controller;
});
