import { helper } from '@ember/component/helper';

export default helper(function mapServiceChecks([checks], hash) {
  //TODO: should namespace be included?
  return checks.map(ck => ck.ID)
});
