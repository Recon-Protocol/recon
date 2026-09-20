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
  return new Intl.NumberFormat('de-DE', {
    minimumFractionDigits: decimals,
    maximumFractionDigits: decimals,
  }).format(value)
}

function formatSupply(value: number | undefined): string {
  if (value === undefined || Number.isNaN(value)) return '—'
  const kas = value / 100_000_000
  if (kas >= 1_000_000_000) return `${formatNumber(kas / 1_000_000_000)} Mrd. KAS`
  if (kas >= 1_000_000) return `${formatNumber(kas / 1_000_000)} Mio. KAS`
  if (kas >= 1_000) return `${formatNumber(kas / 1_000)} Tsd. KAS`
  return `${formatNumber(kas)} KAS`
}

function formatPrice(value: number | undefined): string {
  if (value === undefined || Number.isNaN(value)) return '—'
  return `${value.toLocaleString('de-DE', {
    minimumFractionDigits: 6,
    maximumFractionDigits: 6,
  })} USD`
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
      setError(err instanceof Error ? err.message : 'Verbindung zur RECON API nicht möglich')
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
            <div className="brand-subtitle">NETZWERK-INTELLIGENZSYSTEM</div>
          </div>
        </div>

        <div className="system-status">
          <button
            className="theme-toggle"
            onClick={() => setDarkMode(!darkMode)}
            title="Theme wechseln"
          >
            {darkMode ? '☀' : '☽'}
          </button>
          <span className="status-dot" />
          <span>SYSTEMBETRIEB / KASPA-MAINNET</span>
          <span className="status-divider">/</span>
          <span>{networkOnline ? 'ONLINE' : 'OFFLINE'}</span>
        </div>
      </header>

      <section className="hero-section">
        <div>
          <p className="eyebrow">RECON PROTOCOL / NETWORK INTELLIGENCE</p>
          <h1>KASPA <span>RECON</span></h1>
          <p className="hero-description">
            Kontinuierliche Netzwerkintelligenz für die Kaspa-Proof-of-Work-Infrastruktur.
            Überwachung des Netzwerkzustands, der Versorgungsdynamik, der Hashrate und der Emissionsbedingungen.
          </p>
        </div>

        <div className="hero-status">
          <span className="hero-status-label">NETZWERKSTATUS</span>
          <strong>{networkOnline ? 'ONLINE' : 'OFFLINE'}</strong>
          <span className="hero-status-line" />
          <span className="hero-status-meta">SERVICE: {health?.service ?? 'recon-api'}</span>
          <span className="hero-status-meta">API: v1</span>
        </div>
      </section>

      {error && (
        <div className="error-banner">API CONNECTION ERROR — {error}</div>
      )}

      <section className="metrics-grid">
        <article className="metric-card">
          <div className="metric-header">
            <span>NETZWERK-HASHRATE</span>
            <span className="metric-index">01</span>
          </div>
          <div className="metric-value">
            {hashrate?.hashrate !== undefined ? formatNumber(hashrate.hashrate) : '—'}
            <small>PH/s</small>
          </div>
          <div className="metric-footer">
            <span>LEISTUNGSNACHWEIS</span>
            <span className="positive">LIVE</span>
          </div>
        </article>

        <article className="metric-card">
          <div className="metric-header">
            <span>VORRAT ABGEBAUT</span>
            <span className="metric-index">02</span>
          </div>
          <div className="metric-value">
            {analytics?.percent_mined !== undefined ? `${formatNumber(analytics.percent_mined)}` : '—'}
            <small>%</small>
          </div>
          <div className="metric-footer">
            <span>ZIRKULIEREND</span>
            <span className="positive">AKTIV</span>
          </div>
        </article>

        <article className="metric-card">
          <div className="metric-header">
            <span>BLOCKBELOHNUNG</span>
            <span className="metric-index">03</span>
          </div>
          <div className="metric-value">
            {supply?.block_reward ?? '—'}
            <small>KAS</small>
          </div>
          <div className="metric-footer">
            <span>NÄCHSTE REDUZIERUNG</span>
            <span>
              {analytics?.days_to_reduction !== undefined
                ? `${formatNumber(analytics.days_to_reduction, 1)} TAGE`
                : '—'}
            </span>
          </div>
        </article>

        <article className="metric-card">
          <div className="metric-header">
            <span>INTELLIGENZ</span>
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
              <span className="panel-label">01 / NETZWERK</span>
              <h2>Netzwerkmetriken</h2>
            </div>
            <span className="panel-status">LIVE</span>
          </div>
          <div className="data-list">
            <div className="data-row">
              <span>Netzwerk</span>
              <strong>{metrics?.network ?? 'kaspa-mainnet'}</strong>
            </div>
            <div className="data-row">
              <span>Blockanzahl</span>
              <strong>{metrics ? formatNumber(metrics.block_count, 0) : '—'}</strong>
            </div>
            <div className="data-row">
              <span>Anzahl der Header</span>
              <strong>{metrics ? formatNumber(metrics.header_count, 0) : '—'}</strong>
            </div>
            <div className="data-row">
              <span>Schwierigkeit</span>
              <strong>{metrics ? formatNumber(metrics.difficulty, 0) : '—'}</strong>
            </div>
            <div className="data-row">
              <span>Virtuelle DAA-Punktzahl</span>
              <strong>{metrics ? formatNumber(metrics.virtual_daa_score, 0) : '—'}</strong>
            </div>
            <div className="data-row">
              <span>Hashrate</span>
              <strong>{hashrate?.hashrate !== undefined ? `${formatNumber(hashrate.hashrate)} PH/s` : '—'}</strong>
            </div>
          </div>
        </article>

        <article className="panel">
          <div className="panel-header">
            <div>
              <span className="panel-label">02 / VERSORGUNG</span>
              <h2>Emissionszustand</h2>
            </div>
            <span className="panel-status">SUPPLY</span>
          </div>
          <div className="supply-visual">
            <div className="supply-ring">
              <div>
                <strong>
                  {analytics?.percent_mined !== undefined
                    ? `${formatNumber(analytics.percent_mined)}%`
                    : '—'}
                </strong>
                <span>ABGEBAUT</span>
              </div>
            </div>
            <div className="supply-data">
              <div>
                <span>Umlaufangebot</span>
                <strong>{formatSupply(analytics?.circulating_supply)}</strong>
              </div>
              <div>
                <span>Gesamtangebot</span>
                <strong>{formatSupply(analytics?.max_supply)}</strong>
              </div>
              <div>
                <span>Restbestand</span>
                <strong>{formatSupply(analytics?.remaining_supply)}</strong>
              </div>
              <div>
                <span>Nächste Reduzierung</span>
                <strong>{supply?.next_reduction ?? '—'}</strong>
              </div>
            </div>
          </div>
        </article>

        <article className="panel wide-panel">
          <div className="panel-header">
            <div>
              <span className="panel-label">03 / INTELLIGENZ</span>
              <h2>Netzwerkbewertung</h2>
            </div>
            <span className="panel-status">{intelligence?.trend?.toUpperCase() ?? '—'}</span>
          </div>
          <div className="assessment">
            <div className="score-block">
              <span>AUFKLÄRUNGSWERT</span>
              <strong>{intelligence?.score ?? '—'}</strong>
              <small>/ 100 — {intelligence?.grade ?? '—'}</small>
            </div>
            <div className="signals">
              {(intelligence?.signals ?? []).map((signal, index) => (
                <div className="signal" key={`${signal}-${index}`}>
                  <span>+</span>
                  <div>
                    <strong>{signal}</strong>
                    <small>NETZWERKSIGNAL ERKANNT</small>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </article>

        <article className="panel wide-panel">
          <div className="panel-header">
            <div>
              <span className="panel-label">04 / MARKT</span>
              <h2>Marktdaten</h2>
            </div>
            <span className="panel-status">LIVE</span>
          </div>
          <div className="data-list">
            <div className="data-row">
              <span>KAS-Preis</span>
              <strong>{formatPrice(price?.price_usd)}</strong>
            </div>
            <div className="data-row">
              <span>Marktkapitalisierung</span>
              <strong>
                {price?.market_cap_usd && price.market_cap_usd > 0
                  ? `$${formatNumber(price.market_cap_usd, 0)}`
                  : 'DATEN NICHT VERFÜGBAR'}
              </strong>
            </div>
            <div className="data-row">
              <span>24-Stunden-Volumen</span>
              <strong>
                {price?.volume_24h_usd && price.volume_24h_usd > 0
                  ? `$${formatNumber(price.volume_24h_usd, 0)}`
                  : 'DATEN NICHT VERFÜGBAR'}
              </strong>
            </div>
            <div className="data-row">
              <span>24-Stunden-Änderung</span>
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
              <span>Netzwerk</span>
              <strong>{report?.network ?? '—'}</strong>
            </div>
            <div className="data-row">
              <span>Hashrate</span>
              <strong>{report?.hashrate ?? '—'}</strong>
            </div>
            <div className="data-row">
              <span>Umlaufangebot</span>
              <strong>{report?.circulating_supply ?? '—'}</strong>
            </div>
            <div className="data-row">
              <span>Abgebaut</span>
              <strong>{report?.percent_mined ?? '—'}</strong>
            </div>
            <div className="data-row">
              <span>Blockbelohnung</span>
              <strong>{report?.block_reward ?? '—'}</strong>
            </div>
            <div className="data-row">
              <span>Nächste Reduzierung</span>
              <strong>{report?.next_reduction ?? '—'}</strong>
            </div>
          </div>
        </article>
      </section>

      <footer className="footer">
        <span>RECON-PROTOKOLL / KASPA-NETZWERK-INTELLIGENZ</span>
        <span>AUTOMATISCHE AKTUALISIERUNG / 30 SEKUNDEN</span>
        <span>
          {loading
            ? 'SYNCHRONISIERE...'
            : lastUpdated
              ? `LETZTES UPDATE / ${lastUpdated.toLocaleTimeString('de-DE')}`
              : 'WARTET AUF DATEN'}
        </span>
      </footer>
    </main>
  )
}

export default App