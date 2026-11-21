define(['jquery', 'bootstrap', 'backend', 'table', 'form'], function ($, undefined, Backend, Table, Form) {

    var Controller = {
        index: function () {
            // 初始化表格参数配置
            Table.api.init({
                extend: {
                    index_url: 'tbl/snapshot/index' + location.search,
                    add_url: 'tbl/snapshot/add',
                    edit_url: 'tbl/snapshot/edit',
                    del_url: 'tbl/snapshot/del',
                    multi_url: 'tbl/snapshot/multi',
                    import_url: 'tbl/snapshot/import',
                    table: 'tbl_snapshot',
                }
            });

            var table = $("#table");

            // 初始化表格
            table.bootstrapTable({
                url: $.fn.bootstrapTable.defaults.extend.index_url,
                pk: 'id',
                sortName: 'id',
                columns: [
                    [
                        {checkbox: true},
                        {field: 'id', title: __('Id')},
                        {field: 'works.name', title: __('A_id'), operate: 'LIKE'},
                        {field: 'bworks.name', title: __('B_id'), operate: 'LIKE'},
                        {field: 'snapshot_mod', title: __('Snapshot_mod'), searchList: {"0":__('Snapshot_mod 0'),"1":__('Snapshot_mod 1')}, formatter: Table.api.formatter.normal},
                        {field: 'status', title: __('Status'),searchList: {"0":__('Status 0'),"1":__('Status 1'),"2":__('Status 2')},formatter: Table.api.formatter.status},
                        {field: 'start_time', title: __('Start_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},
                        {field: 'update_time', title: __('Update_time'), operate:'RANGE', addclass:'datetimerange', autocomplete:false, formatter: Table.api.formatter.datetime},

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
