export interface ApiMeta {
  status: string
  service: string
  api_version: string
}

export interface HealthResponse {
  status: string
  service: string
  version: string
}

export interface NetworkInfo extends ApiMeta {
  network: string
  block_count: number
  header_count: number
  difficulty: number
  virtual_daa_score: number
}

export interface NetworkSnapshot extends NetworkInfo {
  timestamp: number
}

export interface SupplyInfo extends ApiMeta {
  circulating_supply: number
  total_supply: number
  block_reward: string
  next_reduction: string
}

export interface HashrateInfo extends ApiMeta {
  hashrate: number
}

export interface PriceInfo extends ApiMeta {
  price_usd: number
  market_cap_usd: number
  volume_24h_usd: number
  change_24h: number
}

export interface Analytics extends ApiMeta {
  hashrate_phs: number
  circulating_supply: number
  max_supply: number
  percent_mined: number
  remaining_supply: number
  block_reward: string
  next_reduction: string
  days_to_reduction: number
}

export interface ExchangeFlow {
  name: string
  balance_kas: number
  change_24h: number
  change_7d: number
  change_30d: number
}

export interface Report extends ApiMeta {
  report_number: string
  date: string
  calendar_week: string
  network: string
  hashrate: string
  daa_score: string
  circulating_supply: string
  percent_mined: string
  block_reward: string
  next_reduction: string
  network_status: string
  exchange_flows: ExchangeFlow[]
  tps: string
  bps: string
  nodes: string
  miner_count: string
  pool_revenue: string
  l2_activity: string
}

export interface Intelligence extends ApiMeta {
  score: number
  grade: string
  trend: string
  signals: string[]
}