import { MetricType } from "../enum/metrics-type";

export interface IMetric {
  type: MetricType;
  path: string;
  headers: string[];
}
