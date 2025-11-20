USE aicongyou;

-- 首先插入一些教师用户数据
INSERT INTO users (name, email, password, role) VALUES
('张教授', 'zhang.professor@university.edu', 'hashed_password_1', 'teacher'),
('李老师', 'li.teacher@university.edu', 'hashed_password_2', 'teacher'),
('王副教授', 'wang.associate@university.edu', 'hashed_password_3', 'teacher'),
('陈讲师', 'chen.lecturer@university.edu', 'hashed_password_4', 'teacher'),
('刘教授', 'liu.professor@university.edu', 'hashed_password_5', 'teacher');

-- 插入cause课程数据
INSERT INTO cause (name, info, semester, teacher_id) VALUES
('计算机科学导论', '本课程介绍计算机科学的基本概念，包括算法、数据结构、编程基础和计算机系统组成。适合计算机专业大一学生。', '2024-9', 1),
('数据结构与算法', '深入学习各种数据结构和算法设计与分析技术，包括数组、链表、树、图等数据结构以及排序、搜索等算法。', '2024-9', 1),
('数据库系统原理', '介绍数据库系统的基本原理、关系模型、SQL语言、数据库设计和事务管理等核心概念。', '2024-9', 2),
('Web开发技术', '涵盖HTML、CSS、JavaScript等前端技术以及Node.js、Express等后端开发技术,实践项目驱动学习。', '2025-2', 3),
('人工智能基础', '介绍人工智能的基本概念、搜索算法、知识表示、机器学习入门和深度学习基础。', '2025-2', 4),
('软件工程', '讲解软件开发生命周期、需求分析、系统设计、软件测试、项目管理等软件工程核心内容。', '2024-9', 2),
('计算机网络', '深入学习网络协议、网络架构、网络安全和分布式系统等计算机网络核心技术。', '2025-2', 5),
('操作系统原理', '介绍操作系统的基本原理，包括进程管理、内存管理、文件系统和设备管理等核心内容。', '2024-9', 1),
('机器学习实践', '通过实际项目学习机器学习算法和应用，包括监督学习、无监督学习和深度学习实践。', '2025-2', 4),
('移动应用开发', '学习Android和iOS移动应用开发技术,包括UI设计、数据存储、网络通信和性能优化。', '2025-2', 3);


