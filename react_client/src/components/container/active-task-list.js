import TaskList from "../presentation/task-list/task-list";
import {connect} from "react-redux";

const mapStateToProps = state => {
    return {
        tasks: state.tasks,
        loading: state.loading,
    };
};


const ActiveTaskList = connect(
    mapStateToProps,
    null,
)(TaskList);


export default ActiveTaskList;