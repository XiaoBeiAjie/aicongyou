CREATE DATABASE IF NOT EXISTS aicongyou;
USE aicongyou;

CREATE TABLE users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    INDEX idx_deleted_at (deleted_at),
    INDEX idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE users 
ADD COLUMN role VARCHAR(20) NOT NULL DEFAULT 'student' COMMENT '用户角色: student, teacher, admin';


CREATE TABLE cause (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '课程名',
    info VARCHAR(512) NOT NULL DEFAULT '暂无课程简介' COMMENT '课程简介',
    semester VARCHAR(256) NOT NULL DEFAULT '2026-9' COMMENT '开设学期',
    teacher_id BIGINT UNSIGNED NOT NULL COMMENT '授课教师id',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    INDEX idx_deleted_at (deleted_at),
    INDEX idx_semster (semester)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE task (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '任务名',
    cause_id BIGINT UNSIGNED NOT NULL COMMENT '所属的课程id',
    core_requirements VARCHAR(512) COMMENT '任务核心要求',
    submited_count INT NOT NULL DEFAULT 0 COMMENT '已提交人数',
    all_count INT NOT NULL DEFAULT 0 COMMENT '课程全部人数',
    view_count INT NOT NULL DEFAULT 0 COMMENT '查看次数',
    discuss_count INT NOT NULL DEFAULT 0 COMMENT '讨论次数',
    excellent_submission_ids VARCHAR(256) COMMENT '优秀作业id',
    start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '任务开始时间',
    expiration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '任务截止时间',
    score_criteria VARCHAR(512) COMMENT '任务给分标准',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    INDEX idx_deleted_at (deleted_at),
    INDEX idx_cause_id (cause_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE submission (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    content VARCHAR(512) NOT NULL COMMENT '提交内容',
    score INT NOT NULL DEFAULT 0 COMMENT '得分',
    student_id BIGINT UNSIGNED NOT NULL COMMENT '学生id',
    task_id BIGINT UNSIGNED NOT NULL COMMENT '任务id',
    is_submited INT NOT NULL DEFAULT 0 COMMENT '是否提交',
    teacher_comment VARCHAR(512) COMMENT '教师评语',
    is_recommend INT NOT NULL DEFAULT 0 COMMENT '是否推荐',

    INDEX idx_task_id (task_id),
    INDEX idx_is_recommend (is_recommend) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;