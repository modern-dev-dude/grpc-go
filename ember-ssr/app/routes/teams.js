import Route from '@ember/routing/route';
import { inject as service } from '@ember/service';
import AuthService from 'shlack/services/auth';

export const ALL_TEAMS = [
  {
    id: 'linkedin',
    name: 'LinkedIn',
    order: 2,
    iconUrl: '/assets/img/linkedin.png',
  },
  {
    id: 'ms',
    name: 'Microsoft',
    order: 3,
    iconUrl: '/assets/img/microsoft.png',
  },
  {
    id: 'avengers',
    name: 'Avengers',
    order: 4,
    iconUrl: '/assets/img/avengers.jpg',
  },
];

export default class TeamsRoute extends Route {
  /**
   *
   * @type {AuthService}
   */
  @service auth;

  async beforeModel(transition) {
    await super.beforeModel(transition);
    if (!this.auth.currentUserId) {
      this.transitionTo('login');
    }
  }

  async model(){ 
    return Promise.resolve(ALL_TEAMS)
  }
}
