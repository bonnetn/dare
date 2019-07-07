import React from 'react';

import {connect} from "react-redux";
import TaskForm from "../presentation/task-form/task-form";
import {addTask} from "../../store/actions";

const AddTaskFormBase = ({addTask}) => {
    const task = {
        name: 'name',
    };
    return <TaskForm task={task} onSubmit={addTask}/>
};

const mapDispatchToProps = (dispatch) => {
    return {
        addTask: (task) => {
            dispatch(addTask(task))
        },
    }
};

const AddTaskForm = connect(
    null,
    mapDispatchToProps,
)(AddTaskFormBase);

export default AddTaskForm;