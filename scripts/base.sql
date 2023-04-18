CREATE DATABASE IF NOT EXISTS go_blog DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
USE go_blog;
CREATE TABLE `blog_article` (
                                `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                `title` varchar(100) DEFAULT '' COMMENT '文章标题',
                                `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
                                `content` longtext COMMENT '文章内容',
                                `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
                                `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                                `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                                `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
                                `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
                                `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
                                `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

CREATE TABLE `blog_category` (
                                 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                 `name` varchar(100) DEFAULT '' COMMENT '分类名称',
                                 `description` varchar(100) DEFAULT '' COMMENT '分类描述',
                                 `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
                                 `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                                 `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                                 `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
                                 `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
                                 `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
                                 `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分类管理';

CREATE TABLE `blog_article_category` (
                                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                         `article_id` int(10) NOT NULL COMMENT '文章编号',
                                         `category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '分类编号',
                                         `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
                                         `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                                         `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                                         `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
                                         `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
                                         `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
                                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章分类关联';