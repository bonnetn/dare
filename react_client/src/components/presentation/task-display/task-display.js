import React from 'react';
import EditTaskForm from "../../container/edit-task-form";
import DeleteTaskForm from "../../container/delete-task-form";

class TaskDisplay extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            show: false,
        };

        this.toggleEdit = this.toggleEdit.bind(this)
    }

    toggleEdit() {
        this.setState((state, props) => ({
            ...state,
            show: !state.show,
        }));
    }

    render() {
        const task = this.props.task;
        if (!this.state.show) {
            return <div>
                <button onClick={this.toggleEdit}>EDIT</button>
                <DeleteTaskForm uuid={task.uuid}/>
                {task.name} : {task.content}
            </div>;
        } else {
            return <span>
            <EditTaskForm task={task}/>
        </span>
        }
    };
}


export default TaskDisplay;