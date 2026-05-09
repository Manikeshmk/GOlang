import React from "react";

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-500 to-purple-600">
      <div className="container mx-auto px-4 py-16">
        <div className="text-center text-white">
          <h1 className="text-5xl font-bold mb-4">Silent Meeting Summarizer</h1>
          <p className="text-xl mb-8">
            AI-powered meeting analysis and insight extraction
          </p>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8 mt-16">
            <div className="bg-white/10 backdrop-blur p-6 rounded-lg">
              <h3 className="text-2xl font-semibold mb-2">
                Real-Time Processing
              </h3>
              <p>Live transcription and analysis as meetings happen</p>
            </div>
            <div className="bg-white/10 backdrop-blur p-6 rounded-lg">
              <h3 className="text-2xl font-semibold mb-2">Smart Extraction</h3>
              <p>Automatically identify tasks, decisions, and action items</p>
            </div>
            <div className="bg-white/10 backdrop-blur p-6 rounded-lg">
              <h3 className="text-2xl font-semibold mb-2">Deep Insights</h3>
              <p>Analyze sentiment, conflicts, and decision confidence</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
