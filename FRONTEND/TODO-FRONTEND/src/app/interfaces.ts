export enum TaskPriority{
    CRITICAL = 0,
    URGENT = 1,
    IMPORTANT = 2,
    NORMAL = 3,
    EVENTUALLY = 4,
}

export interface Task {
    taskID : Number
    taskName : String
    taskText : String
    taskPriority : TaskPriority
    taskDone : Boolean
    taskDeadline : Date
    taskCreatedAt : Date
}

export interface TasksResponse {
    taskList : Task[]
}

export interface LoginInfo {
    username : string
    password : string
}

export interface RegistrationInfo {
    username : string
    email : string
    password : string
}