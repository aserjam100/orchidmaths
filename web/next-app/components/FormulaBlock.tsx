interface FormulaBlockProps {
  code: string
  label?: string
  className?: string
}

export default function FormulaBlock({ code, label, className = '' }: FormulaBlockProps) {
  return (
    <div className={`flex flex-col items-center w-full max-w-full min-w-0 overflow-hidden ${className}`}>
      <pre className="w-full max-w-full min-w-0 bg-transparent border-l-2 border-orchid-blue rounded-none px-3 sm:px-6 py-2 font-mono text-[13px] sm:text-[14px] text-orchid-secondary whitespace-pre-wrap break-words overflow-x-auto leading-relaxed text-left block">
        {code}
      </pre>
      {label && (
        <span className="font-mono text-[13px] text-orchid-muted mt-1 text-center px-2">
          {label}
        </span>
      )}
    </div>
  )
}
