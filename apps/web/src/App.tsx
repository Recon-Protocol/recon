import { useEffect, useState } from 'react'
import './App.css'

import {
  getAnalytics,
  getHashrate,
  getHealth,
  getIntelligence,
  getMetrics,
  getPrice,
  getReport,
  getSupply,
} from './services/api'

import type {
  Analytics,
  HashrateInfo,
  HealthResponse,
  Intelligence,
  NetworkSnapshot,
  PriceInfo,
  Report,
  SupplyInfo,
} from './types/api'

function formatNumber(value: number | undefined, decimals = 2): string {
  if (value === undefined || Number.isNaN(value)) return '—'
  return new Intl.NumberFormat('en-US', {
    minimumFractionDigits: decimals,
    maximumFractionDigits: decimals,
  }).format(value)
}

function formatSupply(value: number | undefined): string {
  if (value === undefined || Number.isNaN(value)) return '—'
  const kas = value / 100_000_000
  if (kas >= 1_000_000_000) return `${formatNumber(kas / 1_000_000_000)}B KAS`
  if (kas >= 1_000_000) return `${formatNumber(kas / 1_000_000)}M KAS`
  if (kas >= 1_000) return `${formatNumber(kas / 1_000)}K KAS`
  return `${formatNumber(kas)} KAS`
}

function formatPrice(value: number | undefined): string {
  if (value === undefined || Number.isNaN(value)) return '—'
  return `$${value.toLocaleString('en-US', {
    minimumFractionDigits: 6,
    maximumFractionDigits: 6,
  })}`
}

function formatBlockReward(value: string | undefined): string {
  if (!value) return '—'
  const num = parseFloat(value)
  if (Number.isNaN(num)) return value
  return num.toFixed(2) + ' KAS'
}

function App() {
  const [health, setHealth] = useState<HealthResponse | null>(null)
  const [metrics, setMetrics] = useState<NetworkSnapshot | null>(null)
  const [supply, setSupply] = useState<SupplyInfo | null>(null)
  const [analytics, setAnalytics] = useState<Analytics | null>(null)
  const [hashrate, setHashrate] = useState<HashrateInfo | null>(null)
  const [price, setPrice] = useState<PriceInfo | null>(null)
  const [report, setReport] = useState<Report | null>(null)
  const [intelligence, setIntelligence] = useState<Intelligence | null>(null)
  const [loading, setLoading] = useState(true)
  const [lastUpdated, setLastUpdated] = useState<Date | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [darkMode, setDarkMode] = useState(true)

  async function loadData() {
    try {
      setError(null)
      const [
        healthData, metricsData, supplyData, analyticsData,
        hashrateData, priceData, reportData, intelligenceData,
      ] = await Promise.all([
        getHealth(), getMetrics(), getSupply(), getAnalytics(),
        getHashrate(), getPrice(), getReport(), getIntelligence(),
      ])
      setHealth(healthData)
      setMetrics(metricsData)
      setSupply(supplyData)
      setAnalytics(analyticsData)
      setHashrate(hashrateData)
      setPrice(priceData)
      setReport(reportData)
      setIntelligence(intelligenceData)
      setLastUpdated(new Date())
    } catch (err) {
      console.error(err)
      setError(err instanceof Error ? err.message : 'Unable to connect to RECON API')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    void loadData()
    const interval = window.setInterval(() => { void loadData() }, 30_000)
    return () => window.clearInterval(interval)
  }, [])

  const networkOnline = health?.status === 'ok'

  return (
    <main className={`recon-app${darkMode ? '' : ' light-mode'}`}>
      <header className="topbar">
        <div className="brand">
          <div className="brand-mark">R</div>
          <div>
            <div className="brand-name">RECON</div>
            <div className="brand-subtitle">NETWORK INTELLIGENCE SYSTEM</div>
          </div>
        </div>

        <div className="system-status">
          <button
            className="theme-toggle"
            onClick={() => setDarkMode(!darkMode)}
            title="Toggle theme"
          >
            {darkMode ? '☀' : '☽'}
          </button>
          <span className="status-dot" />
          <span>SYSTEM ACTIVE / KASPA-MAINNET</span>
          <span className="status-divider">/</span>
          <span>{networkOnline ? 'ONLINE' : 'OFFLINE'}</span>
        </div>
      </header>

      <section className="hero-section">
        <div>
          <p className="eyebrow">RECON PROTOCOL / NETWORK INTELLIGENCE</p>
          <h1>KASPA <span>RECON</span></h1>
          <p className="hero-description">
            Continuous network intelligence for Kaspa proof-of-work infrastructure.
            Monitoring network state, supply dynamics, hashrate, and emission conditions.
          </p>
        </div>

        <div className="hero-status">
          <span className="hero-status-label">NETWORK STATUS</span>
          <strong>{networkOnline ? 'ONLINE' : 'OFFLINE'}</strong>
          <span className="hero-status-line" />
          <span className="hero-status-meta">
            {networkOnline ? 'KASPA MAINNET CONNECTED' : 'CONNECTION LOST'}
          </span>
        </div>
      </section>

      {error && (
        <div className="error-banner">API CONNECTION ERROR — {error}</div>
      )}

      <section className="metrics-grid">
        <article className="metric-card">
          <div className="metric-header">
            <span>NETWORK HASHRATE</span>
            <span className="metric-index">01</span>
          </div>
          <div className="metric-value">
            {hashrate?.hashrate !== undefined ? formatNumber(hashrate.hashrate / 1000) : '—'}
            <small>PH/s</small>
          </div>
          <div className="metric-footer">
            <span>PROOF OF WORK</span>
            <span className="positive">LIVE</span>
          </div>
        </article>

        <article className="metric-card">
          <div className="metric-header">
            <span>SUPPLY MINED</span>
            <span className="metric-index">02</span>
          </div>
          <div className="metric-value">
            {analytics?.percent_mined !== undefined ? formatNumber(analytics.percent_mined) : '—'}
            <small>%</small>
          </div>
          <div className="metric-footer">
            <span>CIRCULATING</span>
            <span className="positive">ACTIVE</span>
          </div>
        </article>

        <article className="metric-card">
          <div className="metric-header">
            <span>BLOCK REWARD</span>
            <span className="metric-index">03</span>
          </div>
          <div className="metric-value">
  {formatBlockReward(supply?.block_reward)}
</div>
          <div className="metric-footer">
            <span>NEXT REDUCTION</span>
            <span className="positive">
              {analytics?.days_to_reduction !== undefined
                ? `${formatNumber(analytics.days_to_reduction, 1)}d`
                : '—'}
            </span>
          </div>
        </article>

        <article className="metric-card">
          <div className="metric-header">
            <span>INTELLIGENCE</span>
            <span className="metric-index">04</span>
          </div>
          <div className="intelligence-value">
            <strong>{intelligence?.score ?? '—'}</strong>
            <span>{intelligence?.grade ?? '—'}</span>
          </div>
          <div className="metric-footer">
            <span>TREND</span>
            <span className="positive">{intelligence?.trend?.toUpperCase() ?? '—'}</span>
          </div>
        </article>
      </section>

      <section className="dashboard-grid">
        <article className="panel">
          <div className="panel-header">
            <div>
              <span className="panel-label">01 / NETWORK</span>
              <h2>Network Metrics</h2>
            </div>
            <span className="panel-status">LIVE</span>
          </div>
          <div className="data-list">
            <div className="data-row">
              <span>Network</span>
              <strong>{metrics?.network ?? 'kaspa-mainnet'}</strong>
            </div>
            <div className="data-row">
              <span>Block Count</span>
              <strong>{metrics ? formatNumber(metrics.block_count, 0) : '—'}</strong>
            </div>
            <div className="data-row">
              <span>Header Count</span>
              <strong>{metrics ? formatNumber(metrics.header_count, 0) : '—'}</strong>
            </div>
            <div className="data-row">
              <span>Difficulty</span>
              <strong>{metrics ? formatNumber(metrics.difficulty, 0) : '—'}</strong>
            </div>
            <div className="data-row">
              <span>Virtual DAA Score</span>
              <strong>{metrics ? formatNumber(metrics.virtual_daa_score, 0) : '—'}</strong>
            </div>
            <div className="data-row">
              <span>Hashrate</span>
              <strong>{hashrate?.hashrate !== undefined ? `${formatNumber(hashrate.hashrate / 1000)} PH/s` : '—'}</strong>
            </div>
          </div>
        </article>

        <article className="panel">
          <div className="panel-header">
            <div>
              <span className="panel-label">02 / SUPPLY</span>
              <h2>Emission State</h2>
            </div>
            <span className="panel-status">LIVE</span>
          </div>
          <div className="supply-visual">
            <div className="supply-ring">
              <div>
                <strong>
                  {analytics?.percent_mined !== undefined
                    ? `${formatNumber(analytics.percent_mined)}%`
                    : '—'}
                </strong>
                <span>MINED</span>
              </div>
            </div>
            <div className="supply-data">
              <div className="supply-row">
                <span>Circulating Supply</span>
                <strong>{formatSupply(analytics?.circulating_supply)}</strong>
              </div>
              <div className="supply-row">
                <span>Max Supply</span>
                <strong>{formatSupply(analytics?.max_supply)}</strong>
              </div>
              <div className="supply-row">
                <span>Remaining Supply</span>
                <strong>{formatSupply(analytics?.remaining_supply)}</strong>
              </div>
              <div className="supply-row">
                <span>Next Reduction</span>
                <strong>{supply?.next_reduction ?? '—'}</strong>
              </div>
            </div>
          </div>
        </article>

        <article className="panel wide-panel">
          <div className="panel-header">
            <div>
              <span className="panel-label">03 / INTELLIGENCE</span>
              <h2>Network Assessment</h2>
            </div>
            <span className="panel-status">{intelligence?.trend?.toUpperCase() ?? '—'}</span>
          </div>
          <div className="assessment">
            <div className="score-block">
              <span>INTELLIGENCE SCORE</span>
              <strong>{intelligence?.score ?? '—'}</strong>
              <small>/ 100 — {intelligence?.grade ?? '—'}</small>
            </div>
            <div className="signals">
              {(intelligence?.signals ?? []).map((signal, index) => (
                <div className="signal" key={`${signal}-${index}`}>
                  <span>+</span>
                  <div>
                    <strong>{signal}</strong>
                    <small>NETWORK SIGNAL DETECTED</small>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </article>

        <article className="panel wide-panel">
          <div className="panel-header">
            <div>
              <span className="panel-label">04 / MARKET</span>
              <h2>Market Data</h2>
            </div>
            <span className="panel-status">LIVE</span>
          </div>
          <div className="data-list">
            <div className="data-row">
              <span>KAS Price</span>
              <strong>{formatPrice(price?.price_usd)}</strong>
            </div>
            {price?.market_cap_usd && price.market_cap_usd > 1000 && (
              <div className="data-row">
                <span>Market Cap</span>
                <strong>${formatNumber(price.market_cap_usd, 0)}</strong>
              </div>
            )}
            {price?.volume_24h_usd && price.volume_24h_usd > 1000 && (
              <div className="data-row">
                <span>24h Volume</span>
                <strong>${formatNumber(price.volume_24h_usd, 0)}</strong>
              </div>
            )}
            <div className="data-row">
              <span>24h Change</span>
              <strong className={(price?.change_24h ?? 0) >= 0 ? 'positive' : ''}>
                {price?.change_24h !== undefined
                  ? `${price.change_24h >= 0 ? '+' : ''}${formatNumber(price.change_24h)}%`
                  : '—'}
              </strong>
            </div>
          </div>
        </article>

        <article className="panel wide-panel">
          <div className="panel-header">
            <div>
              <span className="panel-label">RECON REPORT</span>
              <h2>Network Intelligence Report</h2>
            </div>
            <span className="panel-status">{report?.network_status?.toUpperCase() ?? '—'}</span>
          </div>
          <div className="data-list">
            <div className="data-row">
              <span>Network</span>
              <strong>{report?.network ?? '—'}</strong>
            </div>
            <div className="data-row">
              <span>Hashrate</span>
              <strong>{report?.hashrate ?? '—'}</strong>
            </div>
            <div className="data-row">
              <span>Circulating Supply</span>
              <strong>{report?.circulating_supply ?? '—'}</strong>
            </div>
            <div className="data-row">
              <span>Mined</span>
              <strong>{report?.percent_mined ?? '—'}</strong>
            </div>
            <div className="data-row">
              <span>Block Reward</span>
              <strong>{formatBlockReward(supply?.block_reward)}</strong>
            </div>
            <div className="data-row">
              <span>Next Reduction</span>
              <strong>{report?.next_reduction ?? '—'}</strong>
            </div>
          </div>
        </article>
      </section>

      <footer className="footer">
        <span>RECON PROTOCOL / KASPA NETWORK INTELLIGENCE</span>
        <span>AUTO UPDATE / 30 SECONDS</span>
        <span>
          {loading
            ? 'SYNCING...'
            : lastUpdated
              ? `LAST UPDATE / ${lastUpdated.toLocaleTimeString('en-US')}`
              : 'WAITING FOR DATA'}
        </span>
      </footer>
    </main>
  )
}

export default App