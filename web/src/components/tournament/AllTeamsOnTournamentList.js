import * as React from "react";
import { connect } from "react-redux"
import { Link } from "@reach/router";
import { Card, Table } from "tabler-react";
import PropTypes from "prop-types";
import {
  fetchAllTournamentTeams as fetchAllTournamentTeamsAction
} from "../../actions/tournaments"


const TeamsRows = ({ teams }) => {
    return teams.map((item) => {
        return (
            <Table.Row key={item.team.id}>
                <Table.Col colSpan={2}><Link to={`/app/team/${item.team.id}`}>{item.team.name}</Link></Table.Col>
            </Table.Row>
        )
    })
}

class AllTeamsOnTournamentList extends React.PureComponent {

    componentDidMount() {
        const { tournaments: { activeTournament }, fetchAllTournamentTeamsAction } = this.props;
        fetchAllTournamentTeamsAction(activeTournament.id)
    }

    render() {
        const { tournaments: { activeTournament, allTeamsOnTournament } } = this.props;
        return (
            <Card>
                <Table
                    cards={true}
                    striped={true}
                    responsive={true}
                    className="table-vcenter"
                >
                    <Table.Header>
                        <Table.Row>
                            <Table.ColHeader colSpan={2}>Name</Table.ColHeader>
                        </Table.Row>
                    </Table.Header>
                    <Table.Body>
                        { activeTournament && allTeamsOnTournament &&
                            <TeamsRows
                                teams={allTeamsOnTournament}
                            />
                        }
                    </Table.Body>
                </Table>
            </Card>
        )
    }
}

AllTeamsOnTournamentList.propTypes = {
    fetchAllTournamentTeamsAction: PropTypes.func,
    tournaments: PropTypes.object,
};

const mapStateToProps = state => ({
    tournaments: state.tournaments
});

const mapDispatchToProps = {
  fetchAllTournamentTeamsAction: (tournamentID) => fetchAllTournamentTeamsAction(tournamentID)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AllTeamsOnTournamentList);
