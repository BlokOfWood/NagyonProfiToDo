export enum TaskPriority{
    CRITICAL = 0,
    URGENT = 1,
    IMPORTANT = 2,
    NORMAL = 3,
    EVENTUALLY = 4,
}

export interface TodoItem {
    taskID : number
    taskName : string
    taskText : string
    taskPriority : TaskPriority
    taskDone : boolean
    taskDeadline : Date
    taskCreatedAt : Date
}

export interface TasksResponse {
    taskList : TodoItem[]
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