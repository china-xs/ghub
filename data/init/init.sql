CREATE TABLE `a_account` (
`id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
`username` varchar(48) NOT NULL DEFAULT '' COMMENT '账号名称',
`pwd` varchar(128) NOT NULL DEFAULT '' COMMENT '账号密码',
`email` varchar (128) NOT NULL DEFAULT '' COMMENT '邮箱地址',
`phone` varchar (11) NOT NULL DEFAULT '' COMMENT '联系电话',
`state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '账号状态：{0:"未开启",1:"正常",2:"锁定",3:"debug"}',
`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='账号表';


CREATE TABLE `a_user2role`(
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id` int (11) NOT NULL DEFAULT '0' COMMENT 'account.id',
    `role_id` int(11) NOT NULL DEFAULT '0' COMMENT 'role.id',
    `operate_id` int(11) NOT NULL DEFAULT '0' COMMENT '操作人ID',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='账号角色关联表';


CREATE TABLE `a_role`(
`id` int NOT NULL AUTO_INCREMENT COMMENT '角色ID',
`p_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级ID',
`nodes` varchar (48) NOT NULL DEFAULT '' COMMENT '上级所有节点',
`name` varchar (64) NOT NULL DEFAULT  '' COMMENT '角色名称',
`desc` varchar (256) NOT NULL DEFAULT '' COMMENT '角色描述',
`operate_id` int(11) NOT NULL DEFAULT '0' COMMENT '最后操作人ID',
`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='账号表';


CREATE TABLE `a_rule` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '角色ID',
    `name` varchar (64)  NOT NULL DEFAULT '' COMMENT '节点|路由名称',
    `desc` varchar (256) NOT NULL DEFAULT  '' COMMENT '节点描述',
    `route` varchar (128) NOT NULL DEFAULT '' COMMENT 'http 路由',
    `operate_id` int(11) NOT NULL DEFAULT '0' COMMENT '最后操作人ID',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='账号表';





account // 用户-多角色
role // 角色-
        -
        -
node // 节点



