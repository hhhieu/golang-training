import http from 'k6/http';
import {check, sleep} from 'k6';

export default function() {
  let res = http.get('http://localhost:8080/incre');
    check(res, { 'success login': (r) => r.status === 200 });
    sleep(0.3);
}