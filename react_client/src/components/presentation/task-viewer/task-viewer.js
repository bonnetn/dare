import React from 'react';
import Spinner from "../spinner/spinner";
import ActiveTaskList from "../../container/active-task-list";
import AddTaskForm from "../../container/add-task-form";
import PropTypes from 'prop-types';


const TaskViewer = ({loading}) => {
    if (loading) {
        return <Spinner/>;
    }
    return (
        <div>
            <h1>Tasks:</h1>
            <ActiveTaskList/>
            <AddTaskForm/>
        </div>
    );
};

TaskViewer.propTypes = {
    loading: PropTypes.bool.isRequired,
};



export default TaskViewer;
