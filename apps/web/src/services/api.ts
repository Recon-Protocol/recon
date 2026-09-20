import type {
  Analytics,
  HashrateInfo,
  HealthResponse,
  Intelligence,
  NetworkInfo,
  NetworkSnapshot,
  PriceInfo,
  Report,
  SupplyInfo,
} from '../types/api'

const API_BASE_URL = 'http://localhost:8081/api/v1'

async function fetchApi<T>(endpoint: string): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${endpoint}`)

  if (!response.ok) {
    throw new Error(
      `RECON API request failed: ${response.status} ${response.statusText}`,
    )
  }

  return response.json() as Promise<T>
}

export function getHealth(): Promise<HealthResponse> {
  return fetchApi<HealthResponse>('/health')
}

export function getNetwork(): Promise<NetworkInfo> {
  return fetchApi<NetworkInfo>('/network')
}

export function getMetrics(): Promise<NetworkSnapshot> {
  return fetchApi<NetworkSnapshot>('/metrics')
}

export function getSupply(): Promise<SupplyInfo> {
  return fetchApi<SupplyInfo>('/supply')
}

export function getHashrate(): Promise<HashrateInfo> {
  return fetchApi<HashrateInfo>('/hashrate')
}

export function getPrice(): Promise<PriceInfo> {
  return fetchApi<PriceInfo>('/price')
}

export function getAnalytics(): Promise<Analytics> {
  return fetchApi<Analytics>('/analytics')
}

export function getReport(): Promise<Report> {
  return fetchApi<Report>('/report')
}

export function getIntelligence(): Promise<Intelligence> {
  return fetchApi<Intelligence>('/intelligence')
}