import Route from '@ember/routing/route';

export default class TeamsTeamRoute extends Route {
    async model({teamId}){
        const res = await fetch(`/api/teams/${teamId}`)
        return res.json()
    }
}
