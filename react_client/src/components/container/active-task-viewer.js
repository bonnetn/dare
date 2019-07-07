import {connect} from "react-redux";
import TaskViewer from "../presentation/task-viewer/task-viewer";
import {fetchTasks} from "../../store/actions";


const mapStateToProps = state => {
    return {
        tasks: state.tasks,
        loading: state.loading,
    };
};
const mapDispatchToProps = (dispatch) => {
    dispatch(fetchTasks());
    return {}
};


const ActiveTaskViewer = connect(
    mapStateToProps,
    mapDispatchToProps,
)(TaskViewer);

export default ActiveTaskViewer;

