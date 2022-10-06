export enum TaskPriority{
    CRITICAL = "Critical",
    URGENT = "Urgent",
    IMPORTANT = "Important",
    NORMAL = "Normal",
    EVENTUALLY = "Eventually",
}

export interface TodoItem {
    todoID : number
    name : string
    description : string
    priority : TaskPriority
    done : boolean
    deadline : Date
    createdAt : Date
}

export interface TodoEditor {
    name : string
    description : string
    priority : TaskPriority
    done : boolean
    deadline : Date
}

export interface TasksResponse {
    todoItemList : TodoItem[]
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

export interface SessionIdResponse{
    sessionID : string
}

export interface CheckSessionIdResponse {
  isCorrectSessionID: boolean
}
